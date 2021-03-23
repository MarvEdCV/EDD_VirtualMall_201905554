import { Component, OnInit } from '@angular/core';
import { TiendasService } from 'src/app/servicios/tiendas.service';
import { Tiendas } from 'src/app/modulos/tiendas';
@Component({
  selector: 'app-tiendas',
  templateUrl: './tiendas.component.html',
  styleUrls: ['./tiendas.component.css']
})
export class TiendasComponent implements OnInit {
  //listaTienda: Tiendas[] =[{"Nombre":"Amaya, Rolón y Chavarría Asociados","Descripcion":"Illum accusantium voluptate voluptatem in corrupti dolorem velit et.","Contacto":"976191834","Calificacion":5,"Logo":"https://economipedia.com/wp-content/uploads/2015/10/apple-300x300.png"},{"Nombre":"Aguayo Cepeda S.A.","Descripcion":"Numquam ea est error inventore et porro veritatis.","Contacto":"949 586 354","Calificacion":4,"Logo":"https://i.pinimg.com/originals/3d/0f/0e/3d0f0e8f600627fde858f6c6e668e999.gif"},{"Nombre":"Amaya, Rolón y Chavarría Asociados","Descripcion":"Enim incidunt beatae enim quisquam harum fuga molestiae at.","Contacto":"915 133 743","Calificacion":3,"Logo":"https://as.com/meristation/imagenes/2021/01/18/mexico/1610944753_187605_1610981923_noticia_normal.jpg"},{"Nombre":"Aguayo Cepeda S.A.","Descripcion":"Animi similique quas quam consectetur dolorem.","Contacto":"961.244.645","Calificacion":4,"Logo":"https://graffica.info/wp-content/uploads/2017/01/Kentucky-Fried-Chicken.jpg"},{"Nombre":"Briones Pineda S.A.","Descripcion":"Minus esse nemo eveniet sapiente iste sapiente repudiandae sapiente.","Contacto":"918-764-088","Calificacion":1,"Logo":"https://static-cse.canva.com/blob/211898/17-50-logotipos-que-te-inspiraran.jpg"}]
  listaTienda: Tiendas[]=[]
  mensajeError: string


  constructor(private TiendasService: TiendasService) {
    this.TiendasService.getListaTiendas().subscribe((dataList:any)=>{
      this.listaTienda = dataList.listaTiendas;
      console.log(dataList)   
    },(err)=>{
      this.mensajeError='No se pudo cargar la lista de carnets'
    })
   }

  ngOnInit(): void {
  }

}
