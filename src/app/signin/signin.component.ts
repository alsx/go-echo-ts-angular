import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { URLSearchParams } from '@angular/http';

import { ApiService } from '../api.service';

// https://www.facebook.com/v2.10/dialog/oauth?client_id=329049280815041
//  &response_type=token&granted_scopes=email,public_profile&redirect_uri=http://localhost:4200/fb-signin

@Component({
  selector: 'app-signin',
  templateUrl: './signin.component.html',
  styleUrls: ['./signin.component.css'],
  providers: [ApiService]
})
export class SigninComponent implements OnInit {

  password: string;
  email: string;
  showPassword = false;
  title = 'Enli';
  fbOauthLink = 'https://www.facebook.com/v2.10/dialog/oauth';
  fbOauthParams = {
    client_id: 329049280815041,
    response_type: 'token',
    granted_scopes: 'email,public_profile',
    redirect_uri: '/fb-signin/' // add host:port here
  };

  constructor(protected route: ActivatedRoute, protected router: Router, protected apiService: ApiService) {
    // Set current host
    const full = location.protocol + '//' + location.hostname + (location.port ? ':' + location.port : '');
    this.fbOauthParams.redirect_uri = full + this.fbOauthParams.redirect_uri;
    const params = new URLSearchParams();
    for (const field of Object.keys(this.fbOauthParams)) {
      params.set(field, this.fbOauthParams[field]);
    }
    this.fbOauthLink += '?' + params.toString();
  }

  togglePassword(): void {
    this.showPassword = !this.showPassword;
  }

  async signIn() {
    try {
      const result = await this.apiService.signIn(this.email, this.password);
      localStorage.setItem('token', result['token']);
      this.router.navigate(['']);
    } catch (err) {
      console.error(err);
    }

  }

  goToFacebook() {
    window.location.href = this.fbOauthLink;
  }

  ngOnInit() {
  }

}
