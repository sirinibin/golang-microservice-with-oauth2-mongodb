import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomeComponent }  from './home/home.component';
import { LoginComponent }  from './login/login.component';
import { SignupComponent }  from './signup/signup.component';
import { DocComponent }  from './doc/doc.component';

const routes: Routes = [
  { path: '', component: HomeComponent },
  { path: 'login', component: LoginComponent },
  { path: 'register', component: SignupComponent },
  { path: 'doc', component: DocComponent },

];

@NgModule({
  imports: [ RouterModule.forRoot(routes) ],
  exports: [ RouterModule ],
  declarations: []
})
export class AppRoutingModule { }
