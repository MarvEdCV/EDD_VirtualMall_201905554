import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CargaPedidosComponent } from './carga-pedidos.component';

describe('CargaPedidosComponent', () => {
  let component: CargaPedidosComponent;
  let fixture: ComponentFixture<CargaPedidosComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CargaPedidosComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(CargaPedidosComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
