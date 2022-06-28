import { Component, OnInit } from '@angular/core';
import { MOCKUP } from 'src/app/data/mockData';

@Component({
  selector: 'app-web-admin',
  templateUrl: './web-admin.component.html',
  styleUrls: ['./web-admin.component.css'],
})
export class WebAdminComponent implements OnInit {
  dataList = MOCKUP;
  constructor() {}

  ngOnInit(): void {}
}
