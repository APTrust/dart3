import logo from '../src/img/dart.png';
import {Greet, DashboardShow} from '../wailsjs/go/main/App';

window.addEventListener("load", function(event) {
    load(DashboardShow);
}); 

function load(fn) {
    try {
        fn()
            .then((result) => {
                console.log(result)
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