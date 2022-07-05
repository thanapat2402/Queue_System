import { Component, ElementRef, OnInit, ViewChild } from '@angular/core';
import {
  FormBuilder,
  FormControl,
  FormGroup,
  Validators,
} from '@angular/forms';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { PostQueue } from 'src/app/models/queue';
import { MasterService } from 'src/app/service/master.service';
import { DetailModalComponent } from '../detail-modal/detail-modal.component';
import { WebAdminComponent } from '../web-admin/web-admin.component';

@Component({
  selector: 'app-add-modal',
  templateUrl: './add-modal.component.html',
  styleUrls: ['./add-modal.component.css'],
})
export class AddModalComponent implements OnInit {
  constructor(
    private modalService: NgbModal,
    private service: MasterService,
    private formBuilder: FormBuilder
  ) {}
  @ViewChild('addModal') addview!: ElementRef;
  @ViewChild(DetailModalComponent) viewDetail!: DetailModalComponent;

  tempData: PostQueue = {
    type: '',
    name: '',
    tel: '',
  };
  ngOnInit(): void {}
  queueForm = new FormGroup({
    name: new FormControl('', Validators.required),
    tel: new FormControl('', Validators.required),
    type: new FormControl('', Validators.required),
  });
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

  showData(code: string) {
    this.viewDetail.getQueue(code);
  }
  save() {
    console.log(this.queueForm.value);
    this.service
      .createQueue(this.queueForm.getRawValue())
      .subscribe((result) => {
        console.log(result);
        if (result.message == 'Created') {
          console.log('Hi');
          console.log(result.data.Code);
          alert(`${result.data.Code} has been created`);
        }
        return result.data.Code;
      });
    this.queueForm.setValue({ type: '', tel: '', name: '' });
    this.modalService.dismissAll();
    // let createdCode = '';
    // console.log(code);
    // createdCode = this.createQueue(code);
    // this.modalService.dismissAll();
    // this.viewDetail.getQueue(createdCode);
  }
}
