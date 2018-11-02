import { Component, OnInit } from '@angular/core';
import { GlobalProvider } from "./../globalprovider";

import { User } from '../user';

import { LoginForm } from '../loginform';


import { UserService } from "../user.service";


import { Router} from "@angular/router";


@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {


  loginform: LoginForm  = {
    username: '',
    password: '',
  };

  constructor(public global: GlobalProvider,private userService: UserService,private router:Router) {


    if(!this.global.isGuest){
      this.router.navigate(['doc']);
    }
  }

  ngOnInit() {
  }

  login(loginform:LoginForm ): void {

    console.log("Inside Login Function");

    this.userService.authorize(loginform);

  }

}
