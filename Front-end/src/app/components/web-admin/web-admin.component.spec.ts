import { HttpClientModule } from '@angular/common/http';
import { ComponentFixture, TestBed } from '@angular/core/testing';

import { WebAdminComponent } from './web-admin.component';

describe('WebAdminComponent', () => {
  let component: WebAdminComponent;
  let fixture: ComponentFixture<WebAdminComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [HttpClientModule],
      declarations: [WebAdminComponent],
    }).compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(WebAdminComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
  it('should clear the list', () => {
    expect(component.clearList([1, 2, 3, 4, 5])).toBe();
  });
  it('Should get the locale time string in TH', () => {
    expect(component.getTimeString('August 19, 1975 23:15:30 GMT+00:00')).toBe(
      '06:15:30'
    );
  });
});
