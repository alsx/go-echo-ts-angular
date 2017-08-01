import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { AuthGuard } from './auth.guard';
import { SigninComponent } from './signin/signin.component';
import { SignupComponent } from './signup/signup.component';
import { FbSigninComponent } from './fb-signin/fb-signin.component';
import { ProfileComponent } from './profile/profile.component';

// https://www.facebook.com/v2.10/dialog/oauth?client_id=329049280815041&response_type=token&granted_scopes=email,public_profile&redirect_uri=http://localhost:4200/fb-signin
// http://localhost:4200/fb-signin#access_token=EAAErRMFZCH8EBAJKndKgQQcDCKF0cDhX6ethjAI902aP7uYcqzBlEbDZBaIAkuYzdS9CZCxKRX5XxEPQhYNtHUcNiaCXQDxmYfnqoexChAODwNxoPraDX0qFZAKKv5xZCgWo1VHD9Y8rrK1PWZB6vlfrK3a0qfrt1kZC9KzotzS7cxfmKFHR4WZCqQDB90JLqwQZD&expires_in=5983
// http://localhost:4200/?#access_token=EAAErRMFZCH8EBACMoRbndIQwZAX2kkVOQ7g145CVAs3jZBIeZAvPqOPw24xjZAHdAEB2W1ziZBTF6mxZBjcsMCnvdvNHrT45E8tSt3Kh4q5bUdzSHgYwquh9lOpHkY6YZACwd8QrSXs7quNYOv8q7pjD6MasXZAI4kHeZCrVM5rRwumrwltaVfUblDFs5pI2vX7S4ZD&expires_in=6149
// http://localhost:4200/?code=AQD5Tqv1IP_xKCIbmTFaJuZsrjjrhQtE2ljUtcu3QYhOBkiTel1YZaxJO335JWpR8c6btY3LY_uCFeHeZmqegYHkwarcekQus09MlehlgXQfM0DoRlxPK-akCVScWRvKVu1azkgQr57W5BF8tpnKLcUrs7XTApQk7ubvpwFWZC2qmZwyrgST8yM5ZBmhxRH-LWy4TakxcT7JcxuZ3YnMJ6k6lUpv1w1s8yDWixWCYUVkKI17VHznjYe31GemXQcfnZ3sGlcN6bh6sIfbZyBoXVoPLA7T1Qy_38YSCJRn8TMroaNc9OsXQ0cFd4N0YZeAeNU#_=_
// https://graph.facebook.com/me?access_token=EAAErRMFZCH8EBAJKndKgQQcDCKF0cDhX6ethjAI902aP7uYcqzBlEbDZBaIAkuYzdS9CZCxKRX5XxEPQhYNtHUcNiaCXQDxmYfnqoexChAODwNxoPraDX0qFZAKKv5xZCgWo1VHD9Y8rrK1PWZB6vlfrK3a0qfrt1kZC9KzotzS7cxfmKFHR4WZCqQDB90JLqwQZD

@NgModule({
  declarations: [
    AppComponent,
    SigninComponent,
    SignupComponent,
    FbSigninComponent,
    ProfileComponent
  ],
  imports: [
    AppRoutingModule,
    BrowserModule,
    FormsModule,
    HttpModule
  ],
  providers: [
    AuthGuard
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
