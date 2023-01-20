export namespace application {
	
	export class Response {
	    content: string;
	    modalContent: string;
	    nav: string;
	
	    static createFrom(source: any = {}) {
	        return new Response(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.content = source["content"];
	        this.modalContent = source["modalContent"];
	        this.nav = source["nav"];
	    }
	}

}

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
	export class RemoteRepository {
	    id: string;
	    name: string;
	    url: string;
	    userId: string;
	    apiToken: string;
	    loginExtra: string;
	    pluginId: string;
	    errors: {[key: string]: string};
	
	    static createFrom(source: any = {}) {
	        return new RemoteRepository(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.url = source["url"];
	        this.userId = source["userId"];
	        this.apiToken = source["apiToken"];
	        this.loginExtra = source["loginExtra"];
	        this.pluginId = source["pluginId"];
	        this.errors = source["errors"];
	    }
	}

}

