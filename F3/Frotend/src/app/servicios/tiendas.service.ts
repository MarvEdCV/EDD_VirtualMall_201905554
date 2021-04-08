import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders} from '@angular/common/http';
import {baseURL} from '../ApiUrl/baseURL';
import { Observable, Subject } from 'rxjs';
import { Productos } from '../modulos/productos';

@Injectable({
  providedIn: 'root'
})
export class TiendasService {
  listaProductos: any[] = [];
  constructor(private http: HttpClient) {
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      }), 
    }
   }

   getListaTiendas():Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      }),
    }
    return this.http.get<any>(baseURL+'api/Listatiendas',httpOptions);
   }
   
   getListaProductos(Nombre:any):Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      }),
    } 
    return this.http.get<any>(baseURL+'api/Listaproductos/'+Nombre,httpOptions)
   }

   addProductToCart(producto: any, rutac: any):Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
    };

    return this.http.post<any>(baseURL+"api/CarroCompras/"+rutac,producto,httpOptions)
  }
    CarritoCompras(){
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
    };
    
    return this.http.get<any>(baseURL+"api/ObtenerCarro",httpOptions)
  }
  deleteProductToCart(producto:any):Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
    };

    return this.http.post<any>(baseURL+"api/deleteCarrito"+producto,httpOptions)
  }
   
}
