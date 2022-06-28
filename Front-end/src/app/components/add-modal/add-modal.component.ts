import { Component, ElementRef, OnInit, ViewChild } from '@angular/core';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { PostQueue } from 'src/app/models/queue';
import { MasterService } from 'src/app/service/master.service';
import { WebAdminComponent } from '../web-admin/web-admin.component';

@Component({
  selector: 'app-add-modal',
  templateUrl: './add-modal.component.html',
  styleUrls: ['./add-modal.component.css'],
})
export class AddModalComponent implements OnInit {
  constructor(private modalService: NgbModal, private service: MasterService) {}
  @ViewChild('addModal') addview!: ElementRef;

  tempData: PostQueue = {
    type: '',
  };
  ngOnInit(): void {}

  open() {
    this.modalService
      .open(this.addview, {
        ariaLabelledBy: 'add-modal',
        animation: true,
        centered: true,
      })
      .result.then(
        (result) => {},
        (reason) => {}
      );
  }
  createQueue(type: string) {
    this.tempData.type = type;
    console.log(this.tempData);
    this.service.createQueue(this.tempData).subscribe((result) => {
      console.log(result);
    });
  }
  save(code: string) {
    console.log(code);
    this.createQueue(code);
    this.modalService.dismissAll();
  }
}
