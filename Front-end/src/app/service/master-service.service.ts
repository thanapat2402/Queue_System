import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { HttpResult } from '../models/http';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root',
})
export class MasterServiceService {
  subscribeInstant: any;

  constructor(private http: HttpClient) {}
  createQueue(payload: string): Observable<HttpResult<string>> {
    this.subscribeInstant = this.http.post(
      `${environment.baseApi}/queue`,
      payload
    );
    return this.subscribeInstant;
  }
}
