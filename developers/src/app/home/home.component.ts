import { Component, OnInit } from '@angular/core';
import { GlobalProvider } from "./../globalprovider";

import { Router} from "@angular/router";

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {


  constructor(public global: GlobalProvider,private router:Router) {

    if(!this.global.isGuest){
      this.router.navigate(['doc']);
    }
  }

  ngOnInit() {
  }

}
