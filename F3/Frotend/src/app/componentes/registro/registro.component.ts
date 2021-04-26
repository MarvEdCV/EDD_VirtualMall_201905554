import { Component, OnInit } from '@angular/core';
import { NumberValueAccessor } from '@angular/forms';
import { UsersService } from 'src/app/servicios/users.service';

@Component({
  selector: 'app-registro',
  templateUrl: './registro.component.html',
  styleUrls: ['./registro.component.css']
})
export class RegistroComponent implements OnInit {
  email: string="";
  password: string="";
  nombre: string="";
  dpi: number=0;
  cuenta: string=""

  constructor(public UsuariosService:UsersService) { }
  registraar() {
    const user = {Dpi:this.dpi, Nombre: this.nombre, Correo: this.email, Password: this.password, Cuenta:this.cuenta };
    this.UsuariosService.registrar(user).subscribe(data => console.log(data),err=>console.log(err),()=>console.log("Finish"));
    console.log(user);
  }

  ngOnInit(): void {

  }

}
