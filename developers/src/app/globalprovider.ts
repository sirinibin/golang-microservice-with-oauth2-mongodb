import { Injectable } from '@angular/core';

import { CookieService } from 'ngx-cookie-service';

@Injectable()
export class GlobalProvider {
    isGuest=true;
    username='';
    title="GoLang / GO RESTful API + MongoDb with OAuth2";
    api_endpoint="http://api.go.mongodb.nintriva.net";


    constructor(private cookieService: CookieService) {


    }
}
