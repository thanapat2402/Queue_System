import { HttpClientTestingModule } from '@angular/common/http/testing';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { AddModalComponent } from '../add-modal/add-modal.component';
import { DetailModalComponent } from '../detail-modal/detail-modal.component';
import { WebAdminComponent } from './web-admin.component';

describe('WebAdminComponent', () => {
  let component: WebAdminComponent;
  let fixture: ComponentFixture<WebAdminComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [HttpClientTestingModule, ReactiveFormsModule],
      declarations: [
        WebAdminComponent,
        AddModalComponent,
        DetailModalComponent,
      ],
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
  it('should get Time string in Thai format', () => {
    expect(component.getTimeString('August 19, 1975 23:15:30 GMT+00:00')).toBe(
      '06:15:30'
    );
  });
});
