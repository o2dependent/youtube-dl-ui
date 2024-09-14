export namespace main {
	
	export class Info {
	    author: string;
	    title: string;
	    duration: string;
	    qualityInfo: QualityInfo[];
	    thumbnails: {URL: string,Width: number,Height: number}[];
	
	    static createFrom(source: any = {}) {
	        return new Info(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.author = source["author"];
	        this.title = source["title"];
	        this.duration = source["duration"];
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

