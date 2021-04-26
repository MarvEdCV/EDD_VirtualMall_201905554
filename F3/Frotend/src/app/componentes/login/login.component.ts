import { Component, OnInit } from '@angular/core';
import { Router,CanActivate } from '@angular/router';
import { UsersService } from 'src/app/servicios/users.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})

export class LoginComponent implements OnInit {
  cuenta: string="";
  password: string="";

  constructor(public UsersService:UsersService, private router: Router)  { }
  login() {

    if (this.cuenta=="EDD2021" && this.password=="1234"){
      console.log('Es un administrador'); 
      this.router.navigate(['/Admin']);
    }else
    {
        console.log('Es un usuario'); 
        this.router.navigate(['/Tiendas']);

    }    
    }

  ngOnInit(): void {
  }

}
