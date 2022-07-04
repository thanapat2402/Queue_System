import { Component, ElementRef, OnInit, ViewChild } from '@angular/core';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { DateAndTime, Queue, QueueDetail } from 'src/app/models/queue';
import { MasterService } from 'src/app/service/master.service';
import { WebAdminComponent } from '../web-admin/web-admin.component';

@Component({
  selector: 'app-detail-modal',
  templateUrl: './detail-modal.component.html',
  styleUrls: ['./detail-modal.component.css'],
})
export class DetailModalComponent implements OnInit {
  constructor(private modalService: NgbModal, private service: MasterService) {}
  @ViewChild('detail') viewDetail!: ElementRef;
  @ViewChild(WebAdminComponent) admin!: WebAdminComponent;

  tempData: any;
  queueDetail: QueueDetail = {
    code: '',
    qr: '',
    timestamp: '',
  };

  open() {
    this.modalService
      .open(this.viewDetail, {
        ariaLabelledBy: 'detail-modal',
        animation: true,
        centered: true,
      })
      .result.then(
        (result) => {},
        (reason) => {}
      );
  }
  getQueue(code: string) {
    this.service.getQueue(code).subscribe((result) => {
      console.log(result.data);
      this.tempData = result.data;
    });
    this.open();
    //console.log(this.getDateTime(this.tempData.timeStamp));
  }

  getDateTime(input: string): DateAndTime {
    const date = new Date(input).toLocaleDateString('th');
    const time = new Date(input).toLocaleTimeString('th');

    return { date: date, time: time };
  }

  ngOnInit(): void {}
}
