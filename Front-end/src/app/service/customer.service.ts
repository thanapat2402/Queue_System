import { Injectable } from '@angular/core';
import { environment } from 'src/environments/environment';
import { Observable, Subject, tap } from 'rxjs';
import { HttpClient } from '@angular/common/http';
import { HttpResult } from '../models/http';
import { TotalQueue } from '../models/customer';

@Injectable({
  providedIn: 'root',
})
export class CustomerService {
  private _refreshRequired = new Subject<void>();

  get RefreshRequired() {
    return this._refreshRequired;
  }
  subscribeInstant: any;
  constructor(private http: HttpClient) {}

  getAmountQueue(): Observable<HttpResult<TotalQueue>> {
    const url = `${environment.baseApi}report`;
    this.subscribeInstant = this.http.get(url);
    return this.subscribeInstant;
  }
}
