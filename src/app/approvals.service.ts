import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class ApprovalsService {

  constructor(private httpClient: HttpClient) {
  }

  getApprovals(): Observable<any> {
    return this.httpClient.get(`${ environment.addressUrl }/approval`,
      {
        headers: new HttpHeaders(
          {
            'Content-Type': 'application/json'
          }).set('Access-Control-Allow-Origin', '*')
      });
  }
}
