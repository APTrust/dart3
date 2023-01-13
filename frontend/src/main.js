import * as app from '../wailsjs/go/main/App';
import $ from "jquery"
import "bootstrap";

window.addEventListener("load", function(event) {
    load(app.DashboardShow)
    initNavObserver()
    attachNavEvents()
    initMainContentObserver()
}); 

function load(fn, param) {
    try {
        // 
        // * List requests may have filter params
        // * Show, edit, and delete request will have UUID param
        //   UUID param is a simple string and is bound from data-func-param
        //   HTML attribute when the event listener is attached. See
        //   initMainContentObserver below.
        // * No params for create requests
        // * Update requests will have object params. The object type will
        //   match the request type. E.g. update app setting will take an
        //   AppSetting object. This should be passed as JSON with correct
        //   typing (i.e. numbers and numbers, not strings; same for booleans).
        //
        // 
        if (!param) {
            param = formToJson()
        }
        console.log(`>>> ${fn.name}:  ${param}`)
        fn(param)
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
        let paramString = item.dataset.funcParam
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
                if (functionName.includes("Delete")) {
                    if (confirm("Do you want to delete this item?")) {
                        load(fn, paramString);    
                    }
                } else {
                    load(fn, paramString);
                }
                console.log("Attached " + functionName)
            })
        }
    })
}

function initNavObserver() {
    const callback = function(mutationsList, observer) {
        attachNavEvents()
    }
    const observer = new MutationObserver(callback);
    const navContainer = document.getElementById("nav")
    observer.observe(navContainer, {childList: true, characterData: true})
    return observer
}

function initPopovers() {
    // Attach popover help tips to dynamically added elements
    var popOverSettings = {
        container: 'body',
        trigger: 'hover',
        html: true,
        selector: '[data-toggle="popover"]',
        content: function () {
            return $('#popover-content').html();
        }
    }
    $('body').popover(popOverSettings);
}

function initMainContentObserver() {
    const callback = function(mutationsList, observer) {
        initPopovers()
    }
    const observer = new MutationObserver(callback);
    const mainContainer = document.getElementById("container")
    observer.observe(mainContainer, {childList: true, characterData: true})
    return observer
}

function formToJson() {
    if (document.querySelector('form') == null) {
        return null
    }
    let data = {}
    let form = document.querySelector('form')
    let vals = Object.values(form)
    for (let i=0; i < vals.length; i++) {
        let key = vals[i].name 
        let value = vals[i].value 
        let cast = ""
        if (vals[i].dataset && vals[i].dataset['cast']) {
            cast = vals[i].dataset['cast']
        }
        // If the field says it must be cast to a different type,
        // cast it here. Most values are strings and do not need
        // to be converted. Failure to cast to number/bool will
        // cause a json.Unmarshal error in the Go code. E.g.
        // "item": "true" cannot be Unmarshalled as "item": true.
        switch (cast) {
            case 'int':
                let intVal = parseInt(value, 10)
                if (isNaN(intVal)) {
                    intVal = 0
                } 
                value = intVal
            case 'float': 
                let floatVal = parseInt(value, 10)
                if (isNaN(floatVal)) {
                    floatVal = 0
                } 
                value = floatVal
            case 'bool':
                value = boolValue(value)
        }
        data[key] = value
    }    
    return JSON.stringify(data)
}

function boolValue(val) {
    if (typeof val === 'boolean') {
        return val;
    }
    var lcString = String(val).toLowerCase();
    var trueValues = ['t', 'true', 'yes', '1'];
    var falseValues = ['f', 'false', 'no', '0'];
    var retValue = false;
    if (trueValues.includes(lcString)) {
        retValue = true;
    }
    return retValue;
}