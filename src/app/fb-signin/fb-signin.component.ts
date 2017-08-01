import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { ApiService } from '../api.service';

@Component({
  selector: 'app-fb-signin',
  templateUrl: './fb-signin.component.html',
  styleUrls: ['./fb-signin.component.css'],
  providers: [ApiService]
})
export class FbSigninComponent implements OnInit {

  constructor(protected route: ActivatedRoute, protected router: Router, protected apiService: ApiService) { }

  async saveFbToken(facebookToken: string) {
    try {
      const result = await this.apiService.signupWithFacebookToken(facebookToken);
      localStorage.setItem('token', result['token']);
      this.router.navigate(['']);
    } catch (err) {
      console.log(err);
    }
  }

  getJsonFromUrl() {
    try {
      const query = window.location.href.split('#')[1];
      const result = {};
      query.split('&').forEach(function (part) {
        const item = part.split('=');
        result[item[0]] = decodeURIComponent(item[1]);
      });
      return result;
    } catch (err) {
      console.log(err);
    }
  }

  ngOnInit() {
    const params = this.getJsonFromUrl();
    this.saveFbToken(params['access_token']);
  }
}
