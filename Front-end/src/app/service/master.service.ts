import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { HttpResult } from '../models/http';
import { environment } from 'src/environments/environment';
import { Observable, Subject, tap } from 'rxjs';
import { HttpResponse, PostQueue, Queue } from '../models/queue';

@Injectable({
  providedIn: 'root',
})
export class MasterService {
  private _refreshRequired = new Subject<void>();

  get RefreshRequired() {
    return this._refreshRequired;
  }
  subscribeInstant: any;
  constructor(private http: HttpClient) {}
  getQueue(code: string): Observable<HttpResult<Queue[]>> {
    this.subscribeInstant = this.http.get(`${environment.baseApi}code/${code}`);
    return this.subscribeInstant;
  }
  //getQueues
  getQueues(code?: string): Observable<HttpResult<Queue[]>> {
    console.log(code);
    const url = code
      ? `${environment.baseApi}${code}`
      : `${environment.baseApi}`;
    this.subscribeInstant = this.http.get(url);
    return this.subscribeInstant;
  }
  //CreateQueue
  createQueue(payload: PostQueue): Observable<HttpResult<HttpResponse>> {
    this.subscribeInstant = this.http
      .post(`${environment.baseApi}`, payload)
      .pipe(
        tap(() => {
          this._refreshRequired.next();
        })
      );
    return this.subscribeInstant;
  }
  //deleteQueue
  deleteQueue(code: string): Observable<HttpResult<Queue[]>> {
    this.subscribeInstant = this.http.delete(`${environment.baseApi}${code}`);
    return this.subscribeInstant;
  }
  acceptQueue(code: string): Observable<HttpResult<Queue[]>> {
    this.subscribeInstant = this.http.delete(
      `${environment.baseApi}Accept/${code}`
    );
    return this.subscribeInstant;
  }
}
