import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { CalendarioPedidosComponent } from './componentes/calendario-pedidos/calendario-pedidos.component';
import { CargaInvComponent } from './componentes/carga-inv/carga-inv.component';
import { CargaPedidosComponent } from './componentes/carga-pedidos/carga-pedidos.component';
import { CarritoComponent } from './componentes/carrito/carrito.component';
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
  path: 'Calendario',
  component: CalendarioPedidosComponent
},
{
  path: 'CargaInv',
  component: CargaInvComponent
},
{
  path: 'CargaPedidos',
  component: CargaPedidosComponent
}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
