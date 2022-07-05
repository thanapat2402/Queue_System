import { TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';

import { HttpClientTestingModule } from '@angular/common/http/testing';
import { AppComponent } from './app.component';
import { AddModalComponent } from './components/add-modal/add-modal.component';
import { MasterService } from './service/master.service';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { DetailModalComponent } from './components/detail-modal/detail-modal.component';

describe('AppComponent', () => {
  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [RouterTestingModule, HttpClientTestingModule],
      declarations: [AppComponent],
      // providers: [
      //   AddModalComponent,
      //   { provide: { MasterService, NgbModal } },
      //   DetailModalComponent,
      //   { provide: { MasterService, NgbModal } },
      // ],
    }).compileComponents();
  });

  it('should create the app', () => {
    const fixture = TestBed.createComponent(AppComponent);
    const app = fixture.componentInstance;
    expect(app).toBeTruthy();
  });

  // it(`should have as title 'Queue System'`, () => {
  //   const fixture = TestBed.createComponent(AppComponent);
  //   const app = fixture.componentInstance;
  //   expect(app.title).toEqual('Queue System');
  // });

  // it('should render title', () => {
  //   const fixture = TestBed.createComponent(AppComponent);
  //   fixture.detectChanges();
  //   const compiled = fixture.nativeElement as HTMLElement;
  //   expect(compiled.querySelector('.content span')?.textContent).toContain(
  //     'Front-end app is running!'
  //   );
  // });
});
