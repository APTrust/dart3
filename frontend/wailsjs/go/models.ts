export namespace main {
	
	export class Response {
	    content: string;
	    modalContent: string;
	    nav: string;
	    error: string;
	    flash: string;
	
	    static createFrom(source: any = {}) {
	        return new Response(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.content = source["content"];
	        this.modalContent = source["modalContent"];
	        this.nav = source["nav"];
	        this.error = source["error"];
	        this.flash = source["flash"];
	    }
	}

}

