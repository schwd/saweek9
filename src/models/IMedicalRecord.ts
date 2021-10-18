export interface MedicalRecordInterface {

    ID: string,

    HospitalNumber: string,
   
    PersonalID: string;
   
    NameTitleID: string;
   
    PatientName	: string;
   
    PatientAge: string;
   
    PatientDob: Date | null;

    PatientTel: string;

    HealthInsuranceID: string;

    RegisterDate: Date | null; 
   
   }