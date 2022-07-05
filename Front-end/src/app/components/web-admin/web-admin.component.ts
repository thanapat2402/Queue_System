import { AfterViewChecked, Component, OnInit, ViewChild } from '@angular/core';
import { PostQueue } from 'src/app/models/queue';
import { MasterService } from 'src/app/service/master.service';
import { AddModalComponent } from '../add-modal/add-modal.component';
import { DetailModalComponent } from '../detail-modal/detail-modal.component';

@Component({
  selector: 'app-web-admin',
  templateUrl: './web-admin.component.html',
  styleUrls: ['./web-admin.component.css'],
})
export class WebAdminComponent implements OnInit {
  //dataList = MOCKUP;
  @ViewChild(AddModalComponent) addQueue!: AddModalComponent;
  @ViewChild(DetailModalComponent) viewDetail!: DetailModalComponent;
  dataList: any = [];
  tempData: PostQueue = { type: '', name: '', tel: '' };
  saveResponse: any;
  detail: any;
  constructor(private service: MasterService) {
    this.getQueues();
    this.service.RefreshRequired.subscribe(() => {
      this.getQueues();
    });
  }

  showData(code: string) {
    this.viewDetail.getQueue(code);
  }
  //getQueues
  getQueues(code?: string) {
    console.log(code);
    if (code) {
      this.service.getQueues(code).subscribe((result) => {
        console.log(result);
        this.dataList = result.data;
        console.log(this.dataList);
      });
    } else {
      this.service.getQueues().subscribe((result) => {
        console.log(result);
        this.dataList = result.data;
        console.log(this.dataList);
      });
    }
  }
  deQueue(code: string) {
    if (confirm(`Do you want to delete ${code}?`)) {
      this.service.deleteQueue(code).subscribe((result) => {
        console.log(result);
        if (result.message === 'Deleted') {
          alert(`${code} has been deleted`);
          this.getQueues();
        }
      });
    }
  }
  clearList(list: any) {
    list = [{}];
  }
  clearQueue() {
    if (confirm('Are you sure to clear all queues?')) {
      this.dataList.forEach((item: any) => {
        this.deQueue(item.Code);
      });
    }
  }
  open() {
    this.addQueue.open();
    this.clearList(this.dataList);
    this.getQueues();
  }
  getTimeString(date: string) {
    return new Date(date).toLocaleTimeString('th');
  }
  ngOnInit(): void {
    this.getQueues();
  }
}
