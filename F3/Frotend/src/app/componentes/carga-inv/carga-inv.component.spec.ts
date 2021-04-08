import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CargaInvComponent } from './carga-inv.component';

describe('CargaInvComponent', () => {
  let component: CargaInvComponent;
  let fixture: ComponentFixture<CargaInvComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CargaInvComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(CargaInvComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
