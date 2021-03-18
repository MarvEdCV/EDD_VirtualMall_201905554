import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CalendarioPedidosComponent } from './calendario-pedidos.component';

describe('CalendarioPedidosComponent', () => {
  let component: CalendarioPedidosComponent;
  let fixture: ComponentFixture<CalendarioPedidosComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CalendarioPedidosComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(CalendarioPedidosComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
