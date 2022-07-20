import { AfterViewChecked, Component, OnInit, ViewChild } from '@angular/core';
import { interval, Subscription } from 'rxjs';
import { MOCKUP } from 'src/app/data/mockData';
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
  private updateSubscription!: Subscription;
  dataList: any = [];
  tempData: PostQueue = { type: '', name: '', tel: '' };
  saveResponse: any;
  detail: any;
  @ViewChild(DetailModalComponent) viewDetail!: DetailModalComponent;
  @ViewChild(AddModalComponent) addQueue!: AddModalComponent;
  constructor(private service: MasterService) {
    this.getQueues();
    this.service.RefreshRequired.subscribe((result) => {
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
          console.log(result.data);
          alert(`${code} has been deleted`);
          this.getQueues();
        }
      });
    }
  }
  clearList(list: any) {
    list = [];
  }
  clearQueue() {
    console.log(this.dataList);
    this.dataList.forEach((item: any) => {
      console.log(item.Code);
      this.service.deleteQueue(item.Code).subscribe((result) => {
        console.log(result);
      });
    });
    this.getQueues();
    // this.dataList.forEach((item: any) => {
    //   this.deQueue(item.Code)
    // };
  }
  open() {
    this.addQueue.open();
    this.clearList(this.dataList);
    this.getQueues();
  }
  getTimeString(date: string) {
    return new Date(date).toLocaleTimeString('th');
  }
  getTypeIcon(type: string) {
    if (type === 'A')
      return 'https://cdn-icons-png.flaticon.com/512/32/32438.png';
    else if (type === 'B')
      return 'https://cdn-icons-png.flaticon.com/512/32/32384.png';
    else if (type === 'C')
      return 'https://cdn-icons-png.flaticon.com/512/33/33308.png';
    else if (type === 'D')
      return 'https://cdn-icons-png.flaticon.com/512/32/32441.png';
    else return '/';
  }

  ngOnInit(): void {
    this.updateSubscription = interval(1000).subscribe((val) => this.getQueues);
    // this.getQueues();
  }
}
