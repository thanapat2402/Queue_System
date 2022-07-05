import { HttpClientTestingModule } from '@angular/common/http/testing';
import { TestBed } from '@angular/core/testing';
import { PostQueue } from '../models/queue';

import { MasterService } from './master.service';

describe('MasterService', () => {
  let service: MasterService;
  const mockData: PostQueue = {
    type: 'A',
    name: 'Nop',
    tel: '086-123-1234',
  };

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
    });
    service = TestBed.inject(MasterService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
  it('should create queue', () => {
    expect(service.createQueue(mockData));
  });
  it('should get all queues', () => {
    expect(service.getQueues());
  });
});
