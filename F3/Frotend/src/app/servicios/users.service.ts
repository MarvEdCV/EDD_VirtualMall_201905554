import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, Subject } from 'rxjs';
import { baseURL } from '../ApiUrl/baseURL';

@Injectable({
  providedIn: 'root'
})
export class UsersService {
  private productAddedSource = new Subject<any>();


  productAdded$ = this.productAddedSource.asObservable();
  constructor(private http: HttpClient) {
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      }),
    }
   }


   registrar(user: any):Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
    };

    return this.http.post<any>(baseURL+"api/Registrar",user,httpOptions)
  }
  mandarcomenT(comentario: any):Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
    };

    return this.http.post<any>(baseURL+"api/Comentar",comentario,httpOptions)
  }
}
