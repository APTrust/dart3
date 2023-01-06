import logo from '../src/img/dart.png';
import * as app from '../wailsjs/go/main/App';

window.addEventListener("load", function(event) {
    load(app.DashboardShow)
    initObserver()
    attachNavEvents()
}); 

function load(fn) {
    try {
        // TODO: Load function params here.
        // 
        // * List requests may have filter params
        // * Show, edit, and delete request will have UUID param
        // * No params for create requests
        // * Update requests will have object params. The object type will
        //   match the request type. E.g. update app setting will take an
        //   AppSetting object.
        //
        // 
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
        if (!item.dataset.funcInitialized) {
            item.addEventListener("click", function(e) {
                e.preventDefault()
                e.stopPropagation()
                let fn = app[functionName];
                if (!fn) {
                    alert("Bad function name: " + functionName)
                    return
                }
                item.dataset.funcInitialized = true
                load(fn);
                console.log("Attached " + functionName)
            })
        }
    })
}

function initObserver() {
    const callback = function(mutationsList, observer) {
        attachNavEvents()
    }
    const observer = new MutationObserver(callback);
    const navContainer = document.getElementById("nav")
    observer.observe(navContainer, {childList: true, characterData: true})
    return observer
}