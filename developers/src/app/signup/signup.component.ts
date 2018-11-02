import { Component, OnInit } from '@angular/core';
import { GlobalProvider } from "./../globalprovider";

import { User } from '../user';

import { UserService } from "../user.service";

import { Router} from "@angular/router";


@Component({
  selector: 'app-signup',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.css']
})
export class SignupComponent implements OnInit {

  user: User = {
    name: '',
    username: '',
    email: '',
    password: ''
  };


  constructor(public global: GlobalProvider,private userService: UserService,private router:Router) {

    if(!this.global.isGuest){
      this.router.navigate(['doc']);
    }

  }

  ngOnInit() {
  }

  createUser(user:User): void {

    console.log("Inside Create User Function");

    this.userService.create(user);

  }

}
