import { Injectable } from '@angular/core';

import { User } from './user';

import { Response } from './response';

import { LoginForm } from './loginform';


import { HttpClient,HttpHeaders  } from '@angular/common/http';


import { GlobalProvider } from "./globalprovider";

import { Router} from "@angular/router";


import { CookieService } from 'ngx-cookie-service';


@Injectable()
export class UserService {

  constructor(public global: GlobalProvider,private http: HttpClient,private router:Router,private cookieService: CookieService) { }

  errors={};

  create(user:User): void {

    console.log("Inside Create Function under UserService");

     let body = JSON.stringify(user);

    let httpOptions = {
        headers: new HttpHeaders({
             'Content-Type':  'application/json',
        })
    };



    this.http.post("/v1/register",body,httpOptions)
        .subscribe(

        (response:Response) => {

          this.router.navigate(['login']);
        },
        (err) => {

          this.errors=err.error.errors;

        },
        () => {
          //Completed

        }
    );



  }

  authorize(loginform:LoginForm): void {

    console.log("Inside Login Function under UserService");

    let body = JSON.stringify(loginform);

    let httpOptions = {
      headers: new HttpHeaders({
        'Content-Type':  'application/json',
      })
    };

    this.http.post("/v1/authorize",body,httpOptions)
        .subscribe(

        (response:Response) => {

            console.log(response);
            console.log("Auth Code:"+response.data.authorization_code);
            this.accesstoken(response.data.authorization_code);

         //this
          //this.router.navigate(['login']);
        },
        (err) => {

          this.errors=err.error.errors;
            console.log(this.errors);

        },
        () => {
          //Completed

        }
    );

  }

  accesstoken(authtoken): void {

    console.log("Inside Accesstoken Function under UserService");

    let data={ "authorization_code": authtoken };

    let body = JSON.stringify(data);

    let httpOptions = {
      headers: new HttpHeaders({
        'Content-Type':  'application/json',
      })
    };

    this.http.post("/v1/accesstoken",body,httpOptions)
        .subscribe(

        (response:Response) => {

          console.log(response);

            this.cookieService.set('accesstoken', response.data.access_token,new Date(response.data.expires_at),"/");
            this.global.isGuest=false;

        
            this.userinfo();

          //this.router.navigate(['login']);
        },
        (err) => {


          console.log(err);
          //this.errors=err.error.errors;

        },
        () => {
          //Completed

        }
    );

  }

    userinfo(): void {

        console.log("Inside Userinfo(ME) Function under UserService");

        let at = this.cookieService.get('accesstoken');
        console.log("Access token:"+at);

        this.http.get("/v1/me?access_token="+at)
            .subscribe(

            (response:Response) => {

                console.log(response);

                this.cookieService.set( 'username', response.data.username );
                this.global.username=this.cookieService.get('username');

                this.router.navigate(['doc']);
            },
            (err) => {


                console.log(err);
                //this.errors=err.error.errors;

            },
            () => {
                //Completed

            }
        );

    }



}
