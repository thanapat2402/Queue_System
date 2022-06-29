import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { WebAdminComponent } from './components/web-admin/web-admin.component';
import { WebCustomerComponent } from './components/web-customer/web-customer.component';
import { AddModalComponent } from './components/add-modal/add-modal.component';
import { HttpClientModule } from '@angular/common/http';
import { DetailModalComponent } from './components/detail-modal/detail-modal.component';
import { ReactiveFormsModule } from '@angular/forms';

@NgModule({
  declarations: [
    AppComponent,
    WebAdminComponent,
    WebCustomerComponent,
    AddModalComponent,
    DetailModalComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    NgbModule,
    HttpClientModule,
    ReactiveFormsModule,
  ],
  providers: [],
  bootstrap: [AppComponent],
})
export class AppModule {}
