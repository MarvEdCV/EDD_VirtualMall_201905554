import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { CarritoComponent } from './componentes/carrito/carrito.component';
import { LoginComponent } from './componentes/login/login.component';
import { ProductosComponent } from './componentes/productos/productos.component';
import { RegistroComponent } from './componentes/registro/registro.component';
import { TiendasComponent } from './componentes/tiendas/tiendas.component';

const routes: Routes = [
  {
    path: 'Tiendas',
    component: TiendasComponent,
  },
{
  path: 'Carrito',
  component: CarritoComponent
},
{
  path: 'Productos/:Nombre',
  component: ProductosComponent
},
{
  path: 'Login',
  component: LoginComponent
},
{
  path: 'Registro',
  component: RegistroComponent
}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
