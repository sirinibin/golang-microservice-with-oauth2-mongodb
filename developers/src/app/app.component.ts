import { Component } from '@angular/core';

import { CookieService } from 'ngx-cookie-service';

import { Router} from "@angular/router";

import { GlobalProvider } from "./globalprovider";


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'Expressjs 4.15 RESTful API with OAuth2';
  api_endpoint='api.nodejs.nintriva.net';

  constructor(public global: GlobalProvider,private cookieService: CookieService,private router:Router) {


    if(this.cookieService.check('accesstoken')){
      this.global.isGuest=false;
      this.global.username=this.cookieService.get('username');
      console.log("Username:"+this.global.username);
    }else {
      this.global.isGuest=true;
    }



  }


}
