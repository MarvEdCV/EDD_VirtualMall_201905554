import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { HttpClientJsonpModule, HttpClientModule } from '@angular/common/http'
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { TiendasComponent } from './componentes/tiendas/tiendas.component';
import { CarritoComponent } from './componentes/carrito/carrito.component';
import { CalendarioPedidosComponent } from './componentes/calendario-pedidos/calendario-pedidos.component';
import { CargaInvComponent } from './componentes/carga-inv/carga-inv.component';
import { CargaPedidosComponent } from './componentes/carga-pedidos/carga-pedidos.component';
import { ProductosComponent } from './componentes/productos/productos.component';
@NgModule({
  declarations: [
    AppComponent,
    TiendasComponent,
    CarritoComponent,
    CalendarioPedidosComponent,
    CargaInvComponent,
    CargaPedidosComponent,
    ProductosComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
