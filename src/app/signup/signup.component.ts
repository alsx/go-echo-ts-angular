import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';

import { ApiService } from '../api.service';

@Component({
  selector: 'app-signup',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.css'],
  providers: [ApiService]
})
export class SignupComponent implements OnInit {
  title = 'Enli';
  name: string;
  email: string;
  password: string;
  showPassword = false;

  constructor(protected route: ActivatedRoute, protected router: Router, protected apiService: ApiService) { }

  async signUp() {
    try {
      const result = await this.apiService.signUp(this.name, this.email, this.password);
      localStorage.setItem('token', result['token']);
      this.router.navigate(['']);
    } catch (err) {
      console.error(err);
    }
  }

  togglePassword(): void {
    this.showPassword = !this.showPassword;
  }

  ngOnInit() {
  }

}
