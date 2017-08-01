import { Injectable } from '@angular/core';
import { Http, Headers } from '@angular/http';
import 'rxjs/add/operator/toPromise';

@Injectable()
export class ApiService {

    private headers = new Headers({ 'Content-Type': 'application/json' });

    constructor(private http: Http) { }

    async signIn(email: string, password: string): Promise<{}> {
        const url = '/api/v1/signin';
        const params = {
            email: email,
            password: password
        };

        const response = await this.http
            .post(url, params, { headers: this.headers })
            .toPromise();
        return response.json() as {};
    }

    async signUp(name: string, email: string, password: string): Promise<{}> {
        const url = '/api/v1/signup';
        const params = {
            name: name,
            email: email,
            password: password
        };

        const response = await this.http
            .post(url, params, { headers: this.headers })
            .toPromise();
        return response.json() as {};
    }

    async getUserInfo(): Promise<{}> {
        const url = '/api/v1/user';
        const token = localStorage.getItem('token');
        this.headers.append('Authorization', `Bearer ${token}`);

        const response = await this.http
            .get(url, { headers: this.headers })
            .toPromise();
        return response.json() as {};
    }

    async signupWithFacebookToken(facebookToken: string): Promise<{}> {
        const url = '/api/v1/fb-signup';
        const params = {
            facebookToken: facebookToken
        };

        const response = await this.http
            .post(url, params, { headers: this.headers })
            .toPromise();
        return response.json() as {};
    }
}
