import { HttpClientTestingModule } from '@angular/common/http/testing';
import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DetailModalComponent } from './detail-modal.component';

describe('DetailModalComponent', () => {
  let component: DetailModalComponent;
  let fixture: ComponentFixture<DetailModalComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [DetailModalComponent],
      imports: [HttpClientTestingModule],
    }).compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(DetailModalComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
