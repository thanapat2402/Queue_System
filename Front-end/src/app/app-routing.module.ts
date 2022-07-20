import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { WebAdminComponent } from './components/web-admin/web-admin.component';
import { WebCustomerComponent } from './components/web-customer/web-customer.component';

const routes: Routes = [
  { path: '', redirectTo: 'admin', pathMatch: 'full' },
  { path: 'admin', component: WebAdminComponent },
  { path: 'customer', component: WebCustomerComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
