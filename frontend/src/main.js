import logo from '../src/img/dart.png';
import * as app from '../wailsjs/go/main/App';

window.addEventListener("load", function(event) {
    load(app.DashboardShow);    
    attachNavEvents();
}); 

function load(fn) {
    try {
        fn()
            .then((result) => {
                //console.log(result)
                if (result.content) {
                    document.getElementById("container").innerHTML = result.content;
                }
                if (result.nav) {
                    document.getElementById("nav").innerHTML = result.nav;
                }
                if (result.modalContent) {
                    document.getElementById("modalContent").innerHTML = result.modalContent;
                }                
            })
            .catch((err) => {
                logError(err);
            });
    } catch (err) {
        logError(err);
    }
}

function logError(err) {
    console.log(err)
    try {
        alert(err)
    } catch (ex) {
        console.log(ex)
    }
}


function attachNavEvents() {
    document.querySelectorAll("[data-func]").forEach(function(item){
        let functionName = item.dataset.func;        
        item.addEventListener("click", function(e) {
            e.preventDefault()
            e.stopPropagation()
            let fn = app[functionName];
            if (!fn) {
                alert("Bad function name: " + functionName)
                return
            }
            load(fn);
        })
        console.log("Attached " + functionName)
    })
}

window.attachNavEvents = attachNavEvents;
