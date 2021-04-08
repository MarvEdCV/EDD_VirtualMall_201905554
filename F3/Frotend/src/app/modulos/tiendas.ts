
export class Tiendas {
      Nombre: string 
      Descripcion:  string 
      Contacto:   string 
      Calificacion: number   
      Logo:        string 
  
    constructor(_nombre:string,_Descripcion: string,_Contacto: string,_Calificacion: number,_Logo: string){
      this.Nombre=_nombre
      this.Descripcion=_Descripcion
      this.Contacto=_Contacto
      this.Calificacion = _Calificacion
      this.Logo = _Logo
  } 
   }