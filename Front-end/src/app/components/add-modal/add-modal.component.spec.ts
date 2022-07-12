import { HttpClientModule } from '@angular/common/http';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { PostQueue } from 'src/app/models/queue';

import { AddModalComponent } from './add-modal.component';

const mockData: PostQueue = {
  type: 'A',
  name: 'Golf',
  tel: '0639795144',
};
describe('AddModalComponent', () => {
  let component: AddModalComponent;
  let fixture: ComponentFixture<AddModalComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [HttpClientModule, ReactiveFormsModule],
      declarations: [AddModalComponent],
    }).compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AddModalComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
  it('should required valid name', () => {
    component.queueForm.setValue({
      name: '',
      type: mockData.type,
      tel: mockData.tel,
    });
    expect(component.queueForm.valid).toEqual(false);
  });
  it('should required valid telephone number', () => {
    component.queueForm.setValue({
      name: mockData.name,
      type: mockData.type,
      tel: 'Az120312',
    });
    expect(component.queueForm.valid).toEqual(false);
  });
  it('should required valid type', () => {
    component.queueForm.setValue({
      name: mockData.name,
      type: '',
      tel: mockData.tel,
    });
    expect(component.queueForm.valid).toEqual(false);
  });
});
