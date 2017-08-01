import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
import { ApiService } from '../api.service';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css'],
  providers: [ApiService],
})
export class ProfileComponent implements OnInit {
  user: any;

  constructor(protected route: ActivatedRoute, protected router: Router, protected apiService: ApiService) { }

  ngOnInit() {
    this.getUserInfo();
  }

  logOut() {
    localStorage.removeItem('token');
    this.router.navigate(['signin']);
  }

  async getUserInfo() {
    try {
      this.user = await this.apiService.getUserInfo();
    } catch (err) {
      localStorage.removeItem('token');
      this.router.navigate(['signin']);
    }
  }

}
