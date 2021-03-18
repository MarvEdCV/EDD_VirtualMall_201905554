import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders} from '@angular/common/http';
import {baseURL} from '../ApiUrl/baseURL';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class TiendasService {
   
  private baurl="Listatiendas" 

  constructor(private http: HttpClient) {
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      }), 
    }
   }

   getListaTiendas():Observable<any>{
    let url = `${this.baurl}`;
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      }),
    }
    return this.http.get<any>(baseURL+url,httpOptions);
   }
}
