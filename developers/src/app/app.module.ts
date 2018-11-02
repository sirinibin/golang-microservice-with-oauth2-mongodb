import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { GlobalProvider } from "./globalprovider";
import { CookieService } from 'ngx-cookie-service';
import { UserService } from "./user.service";

import { FormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';

import { AppComponent } from './app.component';
import { AppRoutingModule } from './/app-routing.module';
import { HomeComponent } from './home/home.component';
import { LoginComponent } from './login/login.component';
import { SignupComponent } from './signup/signup.component';
import { DocComponent } from './doc/doc.component';

import { ScrollToModule } from 'ng2-scroll-to-el';





@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    LoginComponent,
    SignupComponent,
    DocComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    HttpClientModule,
    ScrollToModule
  ],
  providers: [
    GlobalProvider,
    UserService,
    CookieService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
