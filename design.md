# DART 3 Design

DART 3 will use Wails, separating the app into two components: a graphical front 
end for defining policies, and a back end for implementing them.

Policies include jobs, BagIt profiles, and workflows. Users can define what they
want to do (bag these files, ship them to some other server) through a 
point-and-click interface. The back end will execute the jobs described by the 
policies.

The front end will be written in JavaScript, HTML and CSS, leveraging the 
considerable work done in DART 2. JavaScript, HTML and CSS are well suited to 
this interactive UI work.

The backend will be written in Go, leveraging the work already done on 
dart-runner. Go is well suited to this system-level work. 

## Front End Responsibilities

The JavaScript/HTML front end will be responsible for producing valid models
that the back end can use to execute jobs. These models include:

* App Settings
* BagIt Profiles
* Upload Targets
* Jobs
* Workflows

Users can customize these models through intutive forms (as in DART 2), with
the front end guiding them through the process and validating their work at
each step. 

All models can be serialized to JSON, which the back end can read to do its
work.

## Back End Responsibilities

The back end will execute tasks based on models a user defined. It will report
back to the front end on the status of those tasks as they unfold. Tasks
include:

* Packaging bags according to BagIt profiles.
* Validating bags according to BagIt profiles.
* Uploading bags to remote servers and/or S3 endpoints.
* Running jobs (package-validate-upload).
* Running workflows (i.e. run the same job over batches of files or folders)

## Front-End/Back-End Communication

The front end and back end do not need to communicate, except:

* When the front end tells the back end to run a job. This requires a simple
"Run" button.
* When the back end reports progress and errors as it runs. After
the user clicks Run, the back end emits JSON. The front end decides how to
render it. (E.g. progress bars, error messages, succcess messages.)

The front end merely saves models (jobs, BagIt profiles, etc.) to a datastore.
When the user clicks Run, the front end tells the back end to run the task
(job, workflow, etc.) with the specified UUID.

## Data Storage

DART 2 stored JSON in plain text files for simplicity. DART 3 will use SQLite
for robustness. This will make it easier to sort through large numbers of models,
and can allow us to store essential metadata like manifests and tag files.

## Testing

DART 2 has an extensive test suite written in jest. DART 3 will need to find a 
replacement for jest that works with Wails. 

Testing will also be separated. Front end tests merely need to test that the
UI behaves as expected and produces valid models. Back end tests just need to
ensure that the back end performs its tasks correctly. We already have a 
decent test suite for dart-runner, which we can extend as necessary.

## Why is DART leaving Electron?

The cost of maintaining and extending a substantial Electron/Node.js project
is simply too high for a small team.

Previous versions of DART were built with Electron and relied on Node.js for
core functionality. Electron and Node introduced a number of problems that 
made the project difficult to maintain and extend. These included:

### Module Bloat

The Electron app had almost 2000 dependencies. Most of these were poorly documented, 
with hard to understand, uncommented code. Dependabot was constantly warning us about 
module updates and security problems. Keeping modules up to date was a huge time suck,
especially when modules introduced breaking changes.

Debugging is difficult when tracing callbacks through many nested layers of 
undocumented and poorly written third-party libraries. This often made simple code
changes painful.

Depending on thousands of poorly documented, poorly written, poorly tested and
poorly maintained libraries means you inherit a mountain of technical debt before
you've even written a line of code. The cost of that debt became too high over
the years. 

We spent an enormous number of man hours just trying to make sure things that worked 
yesterday still work today. All of those hours added zero value to the project, while
preventing our small dev team from doing truly useful work on this and other projects.

Continuing with Node and Electron means we are signing up for more wasted time.

### Electron

Electron itself is heavyweight, using huge amounts of RAM and CPU. Frequent releases 
included breaking changes, especially to the security model. Constantly having to 
"fix" working code to conform to Electron changes was not a good use of our time.

In addition, users wanted a version of DART that would run on servers that had
no windowing system. Electron simply can't do this. 

### Node's Unsuitable Async Model

Node's async programming model fundamentally unsuitable for work that **must** be 
synchronous: packaging tar files, zip files, etc. We made Node behave synchronously 
for core packaging operations using the async module. The resulting callback spaghetti 
code was impossible for most people to understand. (Code that does the same work
synchronously is simple and clear in Go.)

Extensibility was one of DART's core goals, and on this, it failed, due to the
unnecessary complexity of applying asynchronous, evented programming paradigms
to tasks not at all suited to those paradigms. 

