import { Component, Input, OnInit } from '@angular/core';
import { TiendasService } from 'src/app/servicios/tiendas.service';
import { Productos } from 'src/app/modulos/productos';
import { ActivatedRoute } from '@angular/router';
@Component({
  selector: 'app-productos',
  templateUrl: './productos.component.html',
  styleUrls: ['./productos.component.css']
})
export class ProductosComponent implements OnInit {
  listaProductos: Productos[]=[]
  mensajeError: string
  Nombre: string
  RutaProd: string

  constructor(private ruta:ActivatedRoute,private TiendaService: TiendasService) {
    this.Nombre=ruta.snapshot.params.Nombre
    this.CargarProductos()
   }  

  ngOnInit(): void {
  }
  CargarProductos():void{
    this.TiendaService.getListaProductos(this.Nombre).subscribe((dataList:any)=>{
      this.listaProductos = dataList.listaProductos;
      console.log(dataList)
    },(err)=>{
      this.mensajeError='No se pudo cargar la lista de productos'
    })
  }
  
  agregarProducto(producto: any, cantidad: any){
    this.RutaProd=this.Nombre.concat("-")
    this.RutaProd=this.RutaProd.concat(cantidad)
    this.TiendaService.addProductToCart(producto, this.RutaProd).subscribe(
      data=>console.log(data),err=>console.log(err),()=>console.log("Finish"));
  }
}   