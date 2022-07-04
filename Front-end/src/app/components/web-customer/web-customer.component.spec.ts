import { ComponentFixture, TestBed } from '@angular/core/testing';

import { WebCustomerComponent } from './web-customer.component';

describe('WebCustomerComponent', () => {
  let component: WebCustomerComponent;
  let fixture: ComponentFixture<WebCustomerComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [WebCustomerComponent],
    }).compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(WebCustomerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
