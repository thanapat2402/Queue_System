import { HttpClientModule } from '@angular/common/http';
import { TestBed } from '@angular/core/testing';
import { PostQueue } from '../models/queue';

import { MasterService } from './master.service';

const mockInput: PostQueue = {
  type: 'A',
  name: 'Golf',
  tel: '01234567890',
};
describe('MasterService', () => {
  let service: MasterService;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientModule],
    });
    service = TestBed.inject(MasterService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
  // it('should create the queue', () => {
  //   expect(service.createQueue(mockInput)).toContain('Created');
  // });
});
