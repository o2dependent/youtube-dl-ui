export namespace main {
	
	export class Info {
	    author: string;
	    title: string;
	    duration: Date;
	    time: Date;
	    qualityInfo: QualityInfo[];
	    thumbnails: {URL: string,Width: number,Height: number}[];
	
	    static createFrom(source: any = {}) {
	        return new Info(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.author = source["author"];
	        this.title = source["title"];
	        this.duration = new Date(source["duration"]);
	        this.time = new Date(source["time"]);
	        this.qualityInfo = source["qualityInfo"];
	        this.thumbnails = source["thumbnails"];
	    }
	}
	export class QualityInfo {
	    quality: string;
	    audioQuality: string;
	    mimeType: string;
	
	    static createFrom(source: any = {}) {
	        return new QualityInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.quality = source["quality"];
	        this.audioQuality = source["audioQuality"];
	        this.mimeType = source["mimeType"];
	    }
	}

}

