import { Component, OnInit } from '@angular/core';
import { Productos } from 'src/app/modulos/productos';
import { TiendasService } from 'src/app/servicios/tiendas.service';

@Component({
  selector: 'app-carrito',
  templateUrl: './carrito.component.html',
  styleUrls: ['./carrito.component.css']
})
export class CarritoComponent implements OnInit {
  ListaProd: Productos[]=[];
  Cant: number[]=[];
  Precios: number;

  constructor(private TiendaService: TiendasService) {
    this.Seeproducts()
   }

  ngOnInit(): void {
  }

  async Seeproducts(){
    await this.TiendaService.CarritoCompras().subscribe((res)=>{
      this.ListaProd=res
      this.Cant=res.Cantidad
      this.Precios=res.Precio
      console.log("CARRITO PROBANDOOO")
      console.log(res)
    })
  }

}
