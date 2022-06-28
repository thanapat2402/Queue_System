import { Component, ElementRef, OnInit, ViewChild } from '@angular/core';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';

@Component({
  selector: 'app-add-modal',
  templateUrl: './add-modal.component.html',
  styleUrls: ['./add-modal.component.css'],
})
export class AddModalComponent implements OnInit {
  constructor(private modalService: NgbModal) {}
  @ViewChild('content') addview!: ElementRef;
  ngOnInit(): void {}

  open() {
    this.modalService
      .open(this.addview, {
        ariaLabelledBy: 'modal-basic-title',
        animation: true,
        centered: true,
      })
      .result.then(
        (result) => {},
        (reason) => {}
      );
  }
  save(code: string) {
    console.log(code);
    this.modalService.dismissAll();
  }
}
