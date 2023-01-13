export namespace common {
	
	export class AppSetting {
	    id: string;
	    name: string;
	    value: string;
	    help: string;
	    errors: {[key: string]: string};
	    userCanDelete: boolean;
	
	    static createFrom(source: any = {}) {
	        return new AppSetting(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.value = source["value"];
	        this.help = source["help"];
	        this.errors = source["errors"];
	        this.userCanDelete = source["userCanDelete"];
	    }
	}

}

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

