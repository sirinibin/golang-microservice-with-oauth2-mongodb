import { Component, OnInit } from '@angular/core';
import { GlobalProvider } from "./../globalprovider";
import { Router} from "@angular/router";
import { CookieService } from 'ngx-cookie-service';
import { ScrollToModule } from 'ng2-scroll-to-el';


@Component({
  selector: 'app-doc',
  templateUrl: './doc.component.html',
  styleUrls: ['./doc.component.css']
})
export class DocComponent implements OnInit {

  constructor(public global: GlobalProvider,private router:Router,private cookieService: CookieService) {

    if(this.global.isGuest){
      this.router.navigate(['login']);
    }
  }

  ngOnInit() {
  }

  logOut():void {

    this.cookieService.delete('username');
    this.cookieService.delete('accesstoken');
    this.global.isGuest=true;

    this.router.navigate(['']);

  }

  scroll(el) {
    el.scrollIntoView();
  }

}
